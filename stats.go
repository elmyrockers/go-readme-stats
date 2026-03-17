package main

import (
	"os"
	"context"
	"fmt"
	"slices"
	// "github.com/davecgh/go-spew/spew"

	"github.com/google/go-github/v60/github"
	"github.com/ajstarks/svgo"
	"github.com/go-enry/go-enry/v2"
)




//------------------------------------------------------------------------------------------------------
type stats struct {
	client *github.Client
}

func newStats() *stats {
	token := os.Getenv( "GITHUB_TOKEN" )
	client := github.NewClient(nil).WithAuthToken( token )
	return &stats{ client: client }
}

func ( s *stats ) get_repos( visibility string, sort string ) []*github.Repository {
	// Set options
		options := &github.RepositoryListByAuthenticatedUserOptions{
		    Visibility:  visibility,
		    Affiliation: "owner",
		    Sort:        sort, //pushed & full_name
		    ListOptions: github.ListOptions{PerPage: 100},
		}

	// Get list of repos
		ctx := context.Background()
		repos, _, err := s.client.Repositories.ListByAuthenticatedUser(ctx, options)
		if err != nil {
			fmt.Println( err )
		}
	return repos
}


type Language struct {
	Name  string
	Bytes int
	Percentage float64
}

func ( s *stats ) most_used_languages( count int ) []Language  {
	repos := s.get_repos( "public", "full_name" )
	ctx := context.Background()

	// Get list of languages with their size
		languages := map[string]int{}
		totalBytes := 0
		for _, repo := range repos {
			if repo.GetFork() { continue } //skip fork

			owner := repo.GetOwner().GetLogin()
			name := repo.GetName()
			langData, _, _ := s.client.Repositories.ListLanguages(ctx, owner, name)
			for lang, bytes := range langData {
				languages[ lang ] += bytes
				totalBytes += bytes
			}
		}

	// Sort languages with their size (from biggest to smallest)
		sortedLanguages := []Language{}
		for name, bytes := range languages {
			sortedLanguages = append( sortedLanguages,Language{
				Name: name,
				Bytes: bytes,
				Percentage: (float64(bytes) / float64(totalBytes)) * 100,
			})
		}
		slices.SortFunc( sortedLanguages, func(a, b Language) int {
			return b.Bytes - a.Bytes
		})

	// Get Most Used of Languages
		total := len(sortedLanguages)
		if count > total { count = total }
		sortedLanguages = sortedLanguages[: count ]
	return sortedLanguages
}


func get_language_color(lang string) string {
    color := enry.GetColor(lang)
    if color == "" {
        return "#858585" // Fallback gray
    }
    return color
}

func (s *stats) generate_svg(filename string, languages []Language) {
	count := len(languages)
	if count == 0 { return }

	// Create canvas with dynamic calculation
			numRows := (count + 1) / 2
			
			rowHeight := 25
			headerArea := 75 
			
			footerPadding := 14
			
			width := 300
			dynamicHeight := headerArea + (numRows * rowHeight) + footerPadding


			f, _ := os.Create(filename)
			defer f.Close()
			canvas := svg.New(f)
			canvas.Start(width, dynamicHeight, fmt.Sprintf("viewBox=\"0 0 %d %d\" fill=\"none\" role=\"img\"", width, dynamicHeight))

	// Set CSS Styles
			canvas.Style("text/css", `
				.header { font: 600 18px 'Segoe UI', Ubuntu, Sans-Serif; fill: #2f80ed; animation: fadeInAnimation 0.8s ease-in-out forwards; }
				.lang-name { font: 400 11px "Segoe UI", Ubuntu, Sans-Serif; fill: #434d58; }
				.stagger { opacity: 0; animation: fadeInAnimation 0.3s ease-in-out forwards; }
				#rect-mask rect { animation: slideInAnimation 1s ease-in-out forwards; }
				.lang-progress { animation: growWidthAnimation 0.6s ease-in-out forwards; }

				@keyframes slideInAnimation { from { width: 0; } to { width: 250px; } }
				@keyframes growWidthAnimation { from { width: 0; } to { width: 100%; } }
				@keyframes fadeInAnimation { from { opacity: 0; } to { opacity: 1; } }
			`)

	// Set Background Card
		canvas.Rect(0, 0, width-1, dynamicHeight-1, "rx:4.5; fill:#fffefe; stroke:#e4e2e2")

	// Set Title
		canvas.Gtransform("translate(25, 35)")
		canvas.Text(0, 0, "Most Used Languages", "class=\"header\"")
		canvas.Gend()

	// Set Progress Bar
		canvas.Gtransform("translate(25, 55)")
		canvas.Def()
			canvas.Mask("rect-mask", 0, 0, 250, 8)
			canvas.Rect(0, 0, 250, 8, "fill:white; rx:5")
			canvas.MaskEnd()
		canvas.DefEnd()

		currentX := 0.0
		for _, lang := range languages {
			segmentWidth := (lang.Percentage / 100.0) * 250.0
			style := fmt.Sprintf("fill:%s; mask:url(#rect-mask)", get_language_color(lang.Name))
			canvas.Rect(int(currentX), 0, int(segmentWidth)+1, 8, style, "class=\"lang-progress\"")
			currentX += segmentWidth
		}
		canvas.Gend() 

	// Set List of Languages with their percentages (With 2-decimal precision: %.2f)
		canvas.Gtransform("translate(25, 80)")
		for i, lang := range languages {
			col := i % 2 
			row := i / 2
			x, y := col*150, row*rowHeight
			delay := 450 + (i * 100)

			canvas.Gtransform(fmt.Sprintf("translate(%d, %d)", x, y))
			canvas.Group(fmt.Sprintf("class=\"stagger\" style=\"animation-delay: %dms\"", delay))
			canvas.Circle(5, 6, 5, "fill:"+get_language_color(lang.Name))
			
			// Two-decimal place formatting
			canvas.Text(15, 10, fmt.Sprintf("%s %.2f%%", lang.Name, lang.Percentage), "class=\"lang-name\"")
			
			canvas.Gend() 
			canvas.Gend() 
		}
		canvas.Gend() 

	canvas.End()
}