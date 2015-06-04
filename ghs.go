package main

const Version string = "0.0.4"

func main() {
	args, opts := GhsOptionParser()
	query := ""

	for _, arg := range args {
		query += arg
	}

	if opts.Language != "" {
		query += " language:" + opts.Language
	}
	if opts.User != "" {
		query += " user:" + opts.User
	}
	if opts.Repository != "" {
		query += " repo:" + opts.Repository
	}

	repos := SearchRepository(opts.Sort, opts.Order, query, opts.Enterprise)
	PrintRepository(repos)
}
