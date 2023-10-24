package test

import (
	"encoding/json"
	"testing"

	kafka "github.com/opensourceways/kafka-lib/agent"
	"github.com/opensourceways/kafka-lib/mq"

	"github.com/opensourceways/xihe-script/infrastructure/message"
)

func TestMqGame(t *testing.T) {
	err := kafka.Init(
		&kafka.Config{Address: "127.0.0.1:9092"},
		mq.NewLogger(), nil, "", true,
	)
	if err != nil {
		t.Fatal(err)
	}

	defer kafka.Exit()

	data1 := message.MatchMessage{
		CompetitionId: "1",
		UserId:        "1",
		Path:          "昇思AI挑战赛-多类别图像分类/submit_result/s9qfqri3zpc8j2x7_1/result_example_5120-2022-8-8-15-3-16.txt",
		Phase:         "final",
		PlayerId:      "1",
	}

	data2 := message.MatchMessage{
		CompetitionId: "2",
		UserId:        "2",
		Path:          "昇思AI挑战赛-多类别图像分类/submit_result/s9qfqri3zpc8j2x7_1/result_example_5120-2022-8-8-15-3-16.txt",
		Phase:         "final",
		PlayerId:      "2",
	}

	data3 := message.MatchMessage{
		CompetitionId: "3",
		UserId:        "3",
		Path:          "昇思AI挑战赛-艺术家画作风格迁移/submit_result/victor_1/result",
		Phase:         "final",
		PlayerId:      "3",
	}

	data4 := message.MatchMessage{
		CompetitionId: "4",
		UserId:        "4",
		Path:          "昇思AI挑战赛-艺术家画作风格迁移/submit_result/victor_1/result",
		Phase:         "final",
		PlayerId:      "4",
	}

	bys1, err := json.Marshal(data1)
	if err != nil {
		t.Fatal(err)
	}
	bys2, err := json.Marshal(data2)
	if err != nil {
		t.Fatal(err)
	}
	bys3, err := json.Marshal(data3)
	if err != nil {
		t.Fatal(err)
	}

	bys4, err := json.Marshal(data4)
	if err != nil {
		t.Fatal(err)
	}

	err = kafka.Publish("xihe_submission_new", nil, bys1)
	if err != nil {
		t.Fatal(err)
	}

	err = kafka.Publish("xihe_submission_new", nil, bys3)
	if err != nil {
		t.Fatal(err)
	}

	err = kafka.Publish("xihe_submission_new", nil, bys2)
	if err != nil {
		t.Fatal(err)
	}

	err = kafka.Publish("xihe_submission_new", nil, bys4)
	if err != nil {
		t.Fatal(err)
	}
}
