package cmd

import (
	"github.com/Jsoneft/customedConfigGen/intergal"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var configFile string
var contentFile string

var nizCmd = &cobra.Command{
	Use:                    "niz",
	Aliases:                nil,
	SuggestFor:             nil,
	Short:                  "生成 针对固定文本的宏录制 配置文件",
	Long:                   "生成 针对固定文本的宏录制 配置文件",
	Example:                "niz -config=1.pro -text=content.log",
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Hidden:                 false,
	Annotations:            nil,
	Version:                "",
	PersistentPreRun:       nil,
	PersistentPreRunE:      nil,
	PreRun:                 nil,
	PreRunE:                nil,
	Run: func(cmd *cobra.Command, args []string) {

		ctF, err := ioutil.ReadFile(contentFile)
		if err != nil {
			log.Fatalf("contentFile err = %v", err)
		}
		log.Printf("val = %v", ' ')
		out := intergal.Work(ctF)
		log.Printf("out = %v", out)
		//cfF,err := ioutil.ReadFile(configFile)
		//if err != nil {
		//	log.Fatalf("configFile err = %v", err)
		//}

	},
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
	TraverseChildren:           false,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}

func init() {
	nizCmd.Flags().StringVarP(&configFile, "config", "c", "", "配置文件")
	nizCmd.Flags().StringVarP(&contentFile, "content", "C", "", "录制内容")
}
