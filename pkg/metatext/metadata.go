package metatext

import (
	"strings"

	"google.golang.org/grpc/metadata"
)

//MetadataTextMap 基于TextMap模式的metadata方式设置和读取方式
type MetadataTextMap struct {
	metadata.MD //metadata的格式
}

func (m MetadataTextMap) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range m.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m MetadataTextMap) Set(key, val string) {
	key = strings.ToLower(key)
	m.MD[key] = append(m.MD[key], val)
}
