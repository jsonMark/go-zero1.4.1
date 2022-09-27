package model

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/gen"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/util"
	file "github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"strings"
)

var (
	// VarStringSrc describes the source file of sql.
	VarStringSrc string
	// VarStringDir describes the output directory of sql.
	VarStringDir string
	// VarBoolCache describes whether the cache is enabled.
	VarBoolCache bool
	// VarBoolIdea describes whether is idea or not.
	VarBoolIdea bool
	// VarStringURL describes the dsn of the sql.
	VarStringURL string
	// VarStringSliceTable describes tables.
	VarStringSliceTable []string
	// VarStringTable describes a table of sql.
	VarStringTable string
	// VarStringStyle describes the style.
	VarStringStyle string
	// VarStringDatabase describes the database.
	VarStringDatabase string
	// VarStringSchema describes the schema of postgresql.
	VarStringSchema string
	// VarStringHome describes the goctl home.
	VarStringHome string
	// VarStringRemote describes the remote git repository.
	VarStringRemote string
	// VarStringBranch describes the git branch of the repository.
	VarStringBranch string
	// VarBoolStrict describes whether the strict mode is enabled.
	VarBoolStrict bool
)

var errNotMatched = errors.New("sql not matched")

type ddlArg struct {
	src, dir    string
	cfg         *config.Config
	cache, idea bool
	database    string
	strict      bool
}

func MysqlDDL(_ *cobra.Command, _ []string) error {
	//migrationnotes.BeforeCommands(VarStringDir, VarStringStyle)
	src := VarStringSrc
	dir := VarStringDir
	cache := VarBoolCache
	idea := VarBoolIdea
	style := VarStringStyle
	database := VarStringDatabase
	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch
	if len(remote) > 0 {
		repo, _ := file.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}
	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}
	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}

	arg := ddlArg{
		src:      src,
		dir:      dir,
		cfg:      cfg,
		cache:    cache,
		idea:     idea,
		database: database,
		strict:   VarBoolStrict,
	}

	fmt.Println("===============打印=================")
	fmt.Println(arg)
	//return nil
	return fromDDL(arg)
}

func fromDDL(arg ddlArg) error {
	log := console.NewConsole(arg.idea)
	src := strings.TrimSpace(arg.src)
	if len(src) == 0 {
		return errors.New("expected path or path globbing patterns, but nothing found")
	}

	files, err := util.MatchFiles(src)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errNotMatched
	}

	generator, err := gen.NewDefaultGenerator(arg.dir, arg.cfg, gen.WithConsoleOption(log))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = generator.StartFromDDL(file, arg.cache, arg.strict, arg.database)
		if err != nil {
			return err
		}
	}

	return nil
}
