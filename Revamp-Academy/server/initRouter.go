package server

import (
	"codeid.revampacademy/controller/jobhireController"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerManager *jobhireController.ControllerManager) *gin.Engine {
	//set router from gin
	router := gin.Default()

	//Membuat router Endpoint
	jobRoute := router.Group("/jobs")
	{
		//Mockup 1
		jobRoute.GET("", controllerManager.GetJobPostMergeControl)
		//Mockup 1 with search and pagination
		jobRoute.GET("/search", controllerManager.GetJobPostSearch)

		jobRoute.GET("/dumpJobs", controllerManager.GetJobPostControl)

		//Mockup 2
		jobRoute.GET("/view/:id", controllerManager.GetJobPostDetailControl)

		//Mockup 3 -- Create
		jobRoute.POST("/posting/create", controllerManager.JobHireController.CreateJobPostController)

		//Mockup 4 -- Create
		jobRoute.PUT("/posting/update/:id", controllerManager.JobHireController.UpdateJobPostController)
		//test
		// jobRoute.GET("/listJobCategory", controllerManager.GetListCategoryControl)

	}

	masterRoute := router.Group("/masterdata")
	{
		masterRoute.GET("/listaddress", controllerManager.GetListAddressControl)
		masterRoute.GET("/listcity", controllerManager.GetListCityControl)
		masterRoute.POST("/addAddress", controllerManager.CreateAddress)
	}

	return router
}
