package internal

import (
	"context"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"

	log "github.com/sirupsen/logrus"
)

type TerraformRunner struct {
	WorkingDir string
	CMD        *tfexec.Terraform
}

var varFiles []string

func Setup() (*TerraformRunner, error) {
	log.WithField("version", Env.TerraformVersion).Info("installing terraform version")

	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion(Env.TerraformVersion)),

		InstallDir: "/tmp",
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.WithField("version", Env.TerraformVersion).Error("error installing Terraform")
		return nil, err
	}

	workingDir := Env.WorkingDir

	tf, err := tfexec.NewTerraform(workingDir, execPath)

	if err != nil {
		log.WithField("working_dir", Env.WorkingDir).Error("error running NewTerraform")
		return nil, err
	}

	files, err := getTfVarFilesPaths(Env.VarFilesPath)

	if err != nil {
		log.WithField("error", err).Error("failed to list files in the var files path")
	}

	varFiles = files

	return &TerraformRunner{
		WorkingDir: Env.WorkingDir,
		CMD:        tf,
	}, nil
}

func (r *TerraformRunner) Init() error {
	log.Info("initializing terraform module")

	if err := r.CMD.Init(context.Background(), tfexec.Upgrade(true)); err != nil {
		return err
	}

	return nil
}

func (r *TerraformRunner) SelectWorkspace(workspace string) error {
	log.WithField("workspace", workspace).Info("selecting workspace")

	if workspace == "" {
		return nil
	}

	spaces, current, err := r.CMD.WorkspaceList(context.Background())

	if err != nil {
		return err
	}

	// if the current namespace is the same as the desired workspace
	if current == workspace {
		return nil
	}

	if arrayContains(spaces, workspace) {
		if err := r.CMD.WorkspaceSelect(context.Background(), workspace); err != nil {
			return err
		}
	} else {
		if err := r.CMD.WorkspaceNew(context.Background(), workspace); err != nil {
			return err
		}
	}

	return nil
}

func (r *TerraformRunner) Apply(opts ...tfexec.ApplyOption) error {
	log.Info("running terraform apply")

	if err := r.CMD.Apply(context.Background(), opts...); err != nil {
		return err
	}

	return nil
}

func (r *TerraformRunner) Plan(opts ...tfexec.PlanOption) error {
	log.Info("running terraform plan")

	diff, err := r.CMD.Plan(context.Background(), opts...)

	if err != nil {
		return err
	}

	if diff {
		log.Info("plan detected some changes")
	}

	return nil
}

func (r *TerraformRunner) Destroy(opts ...tfexec.DestroyOption) error {
	log.Info("running terraform destroy")

	if err := r.CMD.Destroy(context.Background(), opts...); err != nil {
		return err
	}

	return nil
}

func (r *TerraformRunner) GetOutputs() (map[string][]byte, error) {
	log.Info("retrieving outputs for module")

	outputs, err := r.CMD.Output(context.Background())

	if err != nil {
		return nil, err
	}

	result := map[string][]byte{}

	for key, o := range outputs {
		result[key] = []byte(string(o.Value))
	}

	return result, nil
}

func (r *TerraformRunner) GetPlanOptions() []tfexec.PlanOption {
	opts := []tfexec.PlanOption{}

	for _, path := range varFiles {
		opts = append(opts, tfexec.VarFile(path))
	}

	opts = append(opts, tfexec.Out("/tmp/tfplan"))

	return opts
}

func (r *TerraformRunner) GetApplyOptions() []tfexec.ApplyOption {
	opts := []tfexec.ApplyOption{}

	for _, path := range varFiles {
		opts = append(opts, tfexec.VarFile(path))
	}

	return opts
}

func (r *TerraformRunner) GetDestroyOptions() []tfexec.DestroyOption {
	opts := []tfexec.DestroyOption{}

	for _, path := range varFiles {
		opts = append(opts, tfexec.VarFile(path))
	}

	return opts
}
