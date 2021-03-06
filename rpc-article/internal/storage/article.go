package storage

import (
	"time"
)

type Article struct {
	ID      int       `gorm:"type:bigint(20) auto_increment;primary_key" json:"id"`
	OwnerID int       `gorm:"type:bigint(20);unique_index:owner_title_idx;default:0" json:"ownerID,omitempty"`
	Title   string    `gorm:"type:varchar(128);unique_index:owner_title_idx;not null" json:"title,omitempty"`
	Tags    string    `gorm:"type:varchar(256);default:''" json:"tags,omitempty"`
	Brief   string    `gorm:"type:varchar(60);default:''" json:"brief,omitempty"`
	Content string    `gorm:"type:longtext COLLATE utf8mb4_unicode_520_ci;not null" json:"content,omitempty"`
	CTime   time.Time `gorm:"type:timestamp;column:ctime;not null;default:CURRENT_TIMESTAMP;index:ctime_idx" json:"ctime,omitempty"`
	UTime   time.Time `gorm:"type:timestamp;column:utime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index:utime_idx" json:"utime,omitempty"`
}

const ArticleMappingForElasticsearch = `{
    "settings": {
        "analysis": {
            "tokenizer": {
                "ngram_tokenizer": {
                    "type": "nGram",
                    "min_gram": 1,
                    "max_gram": 10,
                    "token_chars": [
                        "letter",
                        "digit"
                    ]
                }
            },
            "analyzer": {
                "ngram_tokenizer_analyzer": {
                    "type": "custom",
                    "tokenizer": "ngram_tokenizer",
                    "filter": [
                        "lowercase"
                    ]
                }
            }
        },
        "max_ngram_diff": "10"
	},
	"mappings": {
		"properties": {
			"id": {
				"type": "long"
			},
			"title": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			},
			"tags": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			},
			"content": {
				"type": "text",
				"analyzer": "ngram_tokenizer_analyzer",
				"search_analyzer": "standard"
			}
		}
	}
}`
