package main

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const (
	Gemini_1_5_turbo = "gemini-1.5-flash"
	Gemini_2_0_turbo = "gemini-2.0-flash"
)

type Gemini struct {
	Model string
	// sql model
	// id               int
	// arch             string
	// parameters       string
	// context_length   int
	// embedding_length int
}

func (g Gemini) QueryText(question string) []string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel(g.Model)
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
	}

	logUsageMetadata(resp)
	return getResponse(resp)
}

func logUsageMetadata(resp *genai.GenerateContentResponse) {
	log.Printf("UsageMetadata.CachedContentTokenCount:%v UsageMetadata.CandidatesTokenCount:%v UsageMetadata.PromptTokenCount%v UsageMetadata.TotalTokenCount: %v",
		resp.UsageMetadata.CachedContentTokenCount,
		resp.UsageMetadata.CandidatesTokenCount,
		resp.UsageMetadata.PromptTokenCount,
		resp.UsageMetadata.TotalTokenCount,
	)
}

func getResponse(resp *genai.GenerateContentResponse) (answer []string) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if text, ok := part.(genai.Text); ok {
					answer = append(answer, string(text))
				}
			}
		}
	}
	return answer
}

func GetListModels(client *genai.Client, ctx context.Context) (models []string) {
	iter := client.ListModels(ctx)
	for {
		m, err := iter.Next()
		if err != nil {
			break
		}
		models = append(models, m.Name[7:])
	}
	return
}

func main() {
	// taskSql, err := tasksql.NewDB("data.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer taskSql.CloseTaskSQl()

	// taskSql.CreateTableIfNotExist("question")
	// taskSql.CreateTableIfNotExist("answer")

	// // {"net/http", "html", "bufio", "embed"}
	// question := "Write an article about the golang embed package."
	// taskSql.Post("question", question)
	// gem1 := Gemini{
	// 	Model: Gemini_2_0_turbo,
	// }
	// log.Println(gem1.Model)
	// answer := gem1.QueryText(question)
	// taskSql.PostTask("answer", answer[0])
	// log.Println(answer)

}
