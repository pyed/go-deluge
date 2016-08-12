package deluge

type Torrents []*Torrent

type Torrent struct {
	Comment             string    `json:"comment"`
	ActiveTime          float64   `json:"active_time"`
	IsSeed              bool      `json:"is_seed"`
	Hash                string    `json:"hash"`
	UploadPayloadRate   float64   `json:"upload_payload_rate"`
	MoveCompletedPath   string    `json:"move_completed_path"`
	Private             bool      `json:"private"`
	TotalPayloadUpload  float64   `json:"total_payload_upload"`
	Paused              bool      `json:"paused"`
	SeedRank            float64   `json:"seed_rank"`
	SeedingTime         float64   `json:"seeding_time"`
	MaxUploadSlots      float64   `json:"max_upload_slots"`
	PrioritizeFirstLast bool      `json:"prioritize_first_last"`
	DistributedCopies   float64   `json:"distributed_copies"`
	DownloadPayloadRate float64   `json:"download_payload_rate"`
	Message             string    `json:"message"`
	NumPeers            float64   `json:"num_peers"`
	MaxDownloadSpeed    float64   `json:"max_download_speed"`
	MaxConnections      float64   `json:"max_connections"`
	Compact             bool      `json:"compact"`
	Ratio               float64   `json:"ratio"`
	TotalPeers          float64   `json:"total_peers"`
	TotalSize           float64   `json:"total_size"`
	TotalWanted         float64   `json:"total_wanted"`
	State               string    `json:"state"`
	FilePriorities      []float64 `json:"file_priorities"`
	MaxUploadSpeed      float64   `json:"max_upload_speed"`
	RemoveAtRatio       bool      `json:"remove_at_ratio"`
	Tracker             string    `json:"tracker"`
	SavePath            string    `json:"save_path"`
	Progress            float64   `json:"progress"`
	TimeAdded           float64   `json:"time_added"`
	TrackerHost         string    `json:"tracker_host"`
	TotalUploaded       float64   `json:"total_uploaded"`
	Files               []struct {
		Index  float64 `json:"index"`
		Path   string  `json:"path"`
		Offset float64 `json:"offset"`
		Size   float64 `json:"size"`
	} `json:"files"`
	TotalDone       float64 `json:"total_done"`
	NumPieces       float64 `json:"num_pieces"`
	TrackerStatus   string  `json:"tracker_status"`
	TotalSeeds      float64 `json:"total_seeds"`
	MoveOnCompleted bool    `json:"move_on_completed"`
	NextAnnounce    float64 `json:"next_announce"`
	StopAtRatio     bool    `json:"stop_at_ratio"`
	// FileProgress        []float64     `json:"file_progress"`
	MoveCompleted       bool          `json:"move_completed"`
	PieceLength         float64       `json:"piece_length"`
	AllTimeDownload     float64       `json:"all_time_download"`
	MoveOnCompletedPath string        `json:"move_on_completed_path"`
	NumSeeds            float64       `json:"num_seeds"`
	Peers               []interface{} `json:"peers"`
	Name                string        `json:"name"`
	Trackers            []struct {
		SendStats    bool    `json:"send_stats"`
		Fails        float64 `json:"fails"`
		Verified     bool    `json:"verified"`
		URL          string  `json:"url"`
		FailLimit    float64 `json:"fail_limit"`
		CompleteSent bool    `json:"complete_sent"`
		Source       float64 `json:"source"`
		StartSent    bool    `json:"start_sent"`
		Tier         float64 `json:"tier"`
		Updating     bool    `json:"updating"`
	} `json:"trackers"`
	TotalPayloadDownload float64 `json:"total_payload_download"`
	IsAutoManaged        bool    `json:"is_auto_managed"`
	SeedsPeersRatio      float64 `json:"seeds_peers_ratio"`
	Queue                float64 `json:"queue"`
	NumFiles             float64 `json:"num_files"`
	Eta                  float64 `json:"eta"`
	StopRatio            float64 `json:"stop_ratio"`
	IsFinished           bool    `json:"is_finished"`
}
