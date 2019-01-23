package app

import (
	"context"
	pb "github.com/chenzhe84/BaiCloud/metadata-service/proto/app"
)

type Handler struct {
	Repo IRepository
}

func (h *Handler) SaveApp(cxt context.Context, app *pb.App, res *pb.Response) error {
	if err := h.Repo.SaveApp(app); err != nil {
		res.Result = false
		res.Message = err.Error()
	} else {
		res.Result = true
	}
	return nil
}

func (h *Handler) GetApp(cxt context.Context, app *pb.App, res *pb.AppResponse) error {
	if app, err := h.Repo.GetAppById(app.Id); err != nil {
		res.Result = false
		res.Message = err.Error()
	} else {
		res.Result = true
		res.App = app
	}
	return nil
}

func (h *Handler) SaveModule(cxt context.Context, module *pb.Module, res *pb.Response) error {
	if err := h.Repo.SaveModule(module); err != nil {
		res.Result = false
		res.Message = err.Error()
	} else {
		res.Result = true
	}
	return nil
}

func (h *Handler) GetModule(cxt context.Context, module *pb.Module, res *pb.ModuleResponse) error {
	if module, err := h.Repo.GetModuleById(module.Id); err != nil {
		res.Result = false
		res.Message = err.Error()
	} else {
		res.Result = true
		res.Module = module
	}
	return nil
}

func (h *Handler) SaveForm(cxt context.Context, form *pb.Form, res *pb.Response) error {
	if err := h.Repo.SaveForm(form); err != nil {
		res.Result = false
		res.Message = err.Error()
	} else {
		res.Result = true
	}
	return nil
}

func (h *Handler) GetForm(cxt context.Context, form *pb.Form, res *pb.FormResponse) error {
	if form, err := h.Repo.GetFormById(form.Id); err != nil {
		res.Result = false
		res.Message = err.Error()
	} else {
		res.Result = true
		res.Form = form
	}
	return nil
}
