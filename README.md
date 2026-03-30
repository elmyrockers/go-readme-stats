# go-readme-stats
![CI](https://github.com/elmyrockers/go-readme-stats/actions/workflows/generate.yml/badge.svg)
![Last Commit](https://img.shields.io/github/last-commit/elmyrockers/go-readme-stats)
![Go](https://img.shields.io/badge/Go%20version-1.25-blue)

<br>
<div align="center">
	<img src="/coding-passion.webp" alt="Coding Passion" align="center" />
</div>
<br><br>

Personal Go tool using GitHub API to track my coding stats.

I created this Go tool to use the GitHub API to generate stats about my coding activity — my commits, my repositories, my languages — so I can track my progress.

## Purpose

This repository was created for **my personal use** to generate **GitHub README statistics cards** displayed at the **front of my GitHub profile**.  

- Fetches repository data via the **GitHub API**
- Generates **SVG stats cards** using **SVGo**  
- Uses a **CI/CD pipeline** with **GitHub Actions[bot]** to update automatically  
- Optionally uses **Docker** for automation  
- Can also be **triggered manually**  
- Can be **scheduled daily at 00:00(UTC) or 8:00AM (Malaysia Time)** to update the stats automatically.
	<br>**Note:** GitHub typically delays these runs, actual execution usually occurs **around 01:00 UTC (9:00 AM Malaysia Time)**.
- **You can see the stats are updated by GitHub Actions[bot]**

## Technology Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GitHub API](https://img.shields.io/badge/GitHub%20API-181717?style=for-the-badge&logo=github&logoColor=white)
![SVGo](https://img.shields.io/badge/SVGo-FF9900?style=for-the-badge)
![CI/CD](https://img.shields.io/badge/CI%2FCD-AUTO?style=for-the-badge&logo=githubactions&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)


## SVG Stats
<br>
<div align="center">
	<img src="/elmyrockers_stat.svg" alt="Most Used Languages" align="center" />
</div>