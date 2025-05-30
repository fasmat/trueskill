# TrueSkill implementation in Go

[![Build Status](https://img.shields.io/github/actions/workflow/status/fasmat/trueskill/ci.yml)](https://github.com/fasmat/trueskill/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/fasmat/trueskill/graph/badge.svg?token=2WFR1O5B42)](https://codecov.io/gh/fasmat/trueskill)
[![Go Report Card](https://goreportcard.com/badge/github.com/fasmat/trueskill)](https://goreportcard.com/report/github.com/fasmat/trueskill)
[![Go Reference](https://pkg.go.dev/badge/github.com/fasmat/trueskill?status.svg)](https://pkg.go.dev/github.com/fasmat/trueskill?tab=doc)
[![License](https://img.shields.io/github/license/fasmat/trueskill)](./LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/fasmat/trueskill)](https://github.com/fasmat/trueskill/releases/latest)

The TrueSkill ranking system is a skill based ranking system developed at
[Microsoft Research for Xbox Live](https://research.microsoft.com/en-us/projects/trueskill/details.aspx).

It is a flexible ranking system that implements skill tracking of players and
teams based on the outcome of games. The TrueSkill algorithm can be used to
arrange interesting matches (matchmaking) and can be used to recognize and
potentially publish the skills of players and teams tracked.

This package is licensed under the MIT license.
