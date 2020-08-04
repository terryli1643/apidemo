#!/bin/sh

daogenerator -output './' -modelpkg 'gitlab.99safe.org/Shadow/shadow-framework/command/daogenerator/model' -genmodel -dsn 'root:111111@/test?charset=utf8&parseTime=True&loc=Local'
goreturns -w model/*.go
goreturns -w dao/*.go
