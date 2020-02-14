/*******************************************************
 *  File        :   config_test.go.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 10:36 下午
 *  Notes       :
 *******************************************************/
package config

import "testing"

func TestWriteDefaultConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test.1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteDefaultConfig(); (err != nil) != tt.wantErr {
				t.Errorf("WriteDefaultConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfigWriter_WriteConfig(t *testing.T) {
	type fields struct {
		filePath string
		config   config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test.1",
			fields{
				filePath: "test.toml",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfgw := &ConfigWriter{
				filePath: tt.fields.filePath,
				config:   tt.fields.config,
			}
			if err := cfgw.WriteConfig(); (err != nil) != tt.wantErr {
				t.Errorf("WriteConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
