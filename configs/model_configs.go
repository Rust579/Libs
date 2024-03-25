package configs

type Configs struct {
	ServiceKeys struct {
		ServiceAssessment string `json:"service_assessment"`
		ServiceBuilder    string `json:"service_builder"`
		ServiceEdp        string `json:"service_edp"`
		ServiceBackup     string `json:"service_backup"`
		ServiceApps       string `json:"service_apps"`
		ServiceCourses    string `json:"service_courses"`
		ServiceShareplace string `json:"service_shareplace"`
		ServiceModeration string `json:"service_moderation"`
		ServiceRobocode   string `json:"service_robocode"`
		ServiceSupport    string `json:"service_support"`
		ServiceClub       string `json:"service_club"`
		ServiceClubWS     string `json:"service_club_ws"`
		ServiceDigital    string `json:"service_digital"`
		ServicePortal     string `json:"service_portal"`
	} `json:"service_keys"`
	Minio struct {
		Endpoint        string `json:"endpoint"`
		AccessKeyID     string `json:"access_key_id"`
		SecretAccessKey string `json:"secret_access_key"`
	} `json:"minio"`
	OneId struct {
		LocalTcpUrl string `json:"local_tcp_url"`
		DevTcpUrl   string `json:"dev_tcp_url"`
		TestTcpUrl  string `json:"test_tcp_url"`
		ProdTcpUrl  string `json:"prod_tcp_url"`
	}
}
