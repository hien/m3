	"github.com/m3db/m3db/x/xio"
	"github.com/m3db/m3x/context"
	"github.com/m3db/m3x/ident"
			logger.Warnf("host doesn't support HugeTLB, proceeding without it")
	hostID, err := cfg.HostID.Resolve()
	if err != nil {
		logger.Fatalf("could not resolve local host ID: %v", err)
	}

		envCfg, err = cfg.EnvironmentConfig.Configure(environment.ConfigurationParameters{
			HostID: hostID,
		})
	segmentReaderPool := xio.NewSegmentReaderPool(
	var identifierPool ident.Pool
		identifierPool = ident.NewPool(
		identifierPool = ident.NewNativePool(