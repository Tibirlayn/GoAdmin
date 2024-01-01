package handlers
/*
import (
	"fmt"
	"github.com/Tibirlayn/GoAdmin/config"
)

func ConnectAccountDB() (db, error)  {
    db, err := config.AccountConfiguration()
    if err != nil {
        fmt.Println("Error connect router:", err)
        return nil, err
    }
    return db, nil
}

func ConnectBattleDB() {

}

func ConnectBillineDB() {

}

func ConnectGameDB() {

}

func ConnectLogDB() {

}

func ConnectParmDB() {

}

func ConnectStaticsDB() {

}

/*
func ConnectDB() {

	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return 
    }

	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("User:", cfg.User)
	fmt.Println("Passeord:", cfg.Password)
	fmt.Println("FNLAccount:", cfg.FNLAccount)
    fmt.Println("FNLBattle:", cfg.FNLBattle)
    fmt.Println("FNLBilling:", cfg.FNLBilling)
    fmt.Println("FNLGame2155:", cfg.FNLGame2155)
    fmt.Println("FNLLog:", cfg.FNLLog)
    fmt.Println("FNLParm:", cfg.FNLParm)
    fmt.Println("FNLStatistics:", cfg.FNLStatistics)


	//подлючение к БД
    connStringAccount := fmt.Sprintf("server=%s;user id=%s;password=%s;account=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLAccount)
    db_account, err := sql.Open("sqlserver", connStringAccount)
    if err != nil {
        log.Fatal(err)
    }
    defer db_account.Close()

    connStringBattle := fmt.Sprintf("server=%s;user id=%s;password=%s;battle=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLBattle)
    db_battle, err := sql.Open("sqlserver", connStringBattle)
    if err != nil {
        log.Fatal(err)
    }
    defer db_battle.Close()
    
    connStringBilling := fmt.Sprintf("server=%s;user id=%s;password=%s;billing=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLBilling)
    db_billing, err := sql.Open("sqlserver", connStringBilling)
    if err != nil {
        log.Fatal(err)
    }
    defer db_billing.Close()

    connStringGame2155 := fmt.Sprintf("server=%s;user id=%s;password=%s;game2155=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLGame2155)
    db_game2155, err := sql.Open("sqlserver", connStringGame2155)
    if err != nil {
        log.Fatal(err)
    }
    defer db_game2155.Close()

    connStringLog := fmt.Sprintf("server=%s;user id=%s;password=%s;log=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLLog)
    db_log, err := sql.Open("sqlserver", connStringLog)
    if err != nil {
        log.Fatal(err)
    }
    defer db_log.Close()

    connStringParm := fmt.Sprintf("server=%s;user id=%s;password=%s;parm=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLParm)
    db_parm, err := sql.Open("sqlserver", connStringParm)
    if err != nil {
        log.Fatal(err)
    }
    defer db_parm.Close()

    connStringStatistics := fmt.Sprintf("server=%s;user id=%s;password=%s;statistics=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLStatistics)
    db_statistics, err := sql.Open("sqlserver", connStringStatistics)
    if err != nil {
        log.Fatal(err)
    }
    defer db_statistics.Close()


	err = db_account.Ping()
    err = db_battle.Ping()
    err = db_billing.Ping()
    err = db_game2155.Ping()
    err = db_log.Ping()
    err = db_parm.Ping()
    err = db_statistics.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
    } else {
        fmt.Println("Успешное подключение к базе данных")
    }

}

*/