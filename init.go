package configure

import core "github.com/procyon-projects/procyon-core"

func init() {
	/* server */
	core.Register(newServerProperties)
	/* data source */
	core.Register(newDataSourceProperties)
}
