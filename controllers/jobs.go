package controllers

import (
	"mail-backend/dao"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)


//InsertJob handler for inserting job into jobs collection.
func InsertJob(c echo.Context) (err error) {
    var ctx = c.Request().Context()
    body := new(models.JobCreate)
    if err = verr.JSONValidate(c, body); err != nil {
        return
    }
    var job = body.Insert()
    if err = dao.InsertJob(ctx, job); err != nil {
        return
    }
    return c.JSON(vmod.RespCreated(job, "mail_job"))
}

//GetJob handler for get job by id
func GetJob(c echo.Context) (err error) {
    var ctx = c.Request().Context()
    //TODO: add filter from query
    var filter = bson.M{"_id": c.Param("id")}
    var job *models.Job
    if job, err = dao.GetJob(ctx, filter); err != nil {
        return
    }
    return c.JSON(vmod.RespSelected(job, "mail_job"))
}

//ListJob handler for list jobs by query filter
func ListJob(c echo.Context) (err error) {
    var ctx = c.Request().Context()
    var filter = bson.M{}
    var jobs []models.Job
    if jobs, err = dao.ListJobs(ctx, filter); err != nil {
        return
    }
    return c.JSON(vmod.RespSelected(jobs, "mail_job"))
}

//UpdateJob handler for update a job in database.
func UpdateJob(c echo.Context) (err error) {
    var ctx = c.Request().Context()
    var body *models.Job
    if err = verr.JSONValidate(c, body); err != nil {
        return
    }
    if err = dao.UpdateJob(ctx, body); err != nil {
        return
    }
    return c.JSON(vmod.RespUpdated("mail_job"))
}
