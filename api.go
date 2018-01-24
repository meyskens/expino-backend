package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getNewsHandler(c echo.Context) error {
	i, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	news, err := getNewsItem(int(i))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, news)
}

func getAllNewsHandler(c echo.Context) error {
	news, err := getNewsItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, news)
}

func addNewsHandler(c echo.Context) error {
	item := NewsItem{}
	c.Bind(&item)
	err := addNewsItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	go sendUpdate()
	return c.JSON(http.StatusOK, item)
}

func editNewsHandler(c echo.Context) error {
	item := NewsItem{}
	c.Bind(&item)
	err := editNewsItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	go sendUpdate()
	return c.JSON(http.StatusOK, item)
}

func deleteNewsHandler(c echo.Context) error {
	i, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	err = deleteNewsItem(int(i))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	go sendUpdate()
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
