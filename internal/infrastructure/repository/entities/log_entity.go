package entities

type LogEntity struct {
	ID        int64
	Level     string
	Timestamp string
	Source    string
	Detail    string
}

func (p *LogEntity) GetCreateQuery() (query string, args []interface{}) {
	query = `INSERT INTO application_logs (
		level,
		source,
		detail
	) VALUES (?,?,?);`
	args = []interface{}{
		p.Level,
		p.Source,
		p.Detail,
	}

	return
}
