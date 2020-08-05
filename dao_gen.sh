#!/bin/sh

daogenerator -output 'domain/'\
 -modelpkg 'github.com/terryli1643/apidemo/domain/model'\
 -dsn 'root:111111@/shadow?charset=utf8&parseTime=True&loc=Local'\
 -filter 'casbin_rule,id_space'\
 -tablePrefix 't_'
goreturns -w domain/model/*.go
goreturns -w domain/dao/*.go
