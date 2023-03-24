package Config

type Settings struct{
	Port 				string 				`json:"port"`
	ConnectionString 	string 				`json:"postgr_conn_string"`
	Env 			 	string				`json:"env"`
}
