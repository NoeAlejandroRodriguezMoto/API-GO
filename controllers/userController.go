package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/go-pg/pg/v10"
    "golang.org/x/crypto/bcrypt"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/models"
)

func Register(c *gin.Context, db *pg.DB) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }

    user.Password = string(hashedPassword)

    _, err = db.Model(&user).Insert()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}