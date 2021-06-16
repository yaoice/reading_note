/*
使用sync包提供的errgroup, errgroup提供同步，error收集，可使用context取消正处于任务中运行的goroutines.
errgroup可直接看源码，源码比较简短🕐
*/
package main

func main() {
	/*
	ctx, cancel := context.WithTimeout(context.Background(), thresholdTime)
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	if felixLive {
		g.Go(func() error {
			if err := checkFelixHealth(ctx, felixLivenessEp, "liveness"); err != nil {
				return fmt.Errorf("calico/node is not ready: Felix is not live: %+v", err)
			}
			return nil
		})
	}

	if birdLive {
		g.Go(func() error {
			if err := checkServiceIsLive([]string{"confd", "bird"}); err != nil {
				return fmt.Errorf("calico/node is not ready: bird/confd is not live: %+v", err)
			}
			return nil
		})
	}

	if bird6Live {
		g.Go(func() error {
			if err := checkServiceIsLive([]string{"confd", "bird6"}); err != nil {
				return fmt.Errorf("calico/node is not ready: bird6/confd is not live: %+v", err)
			}
			return nil
		})
	}

	if felixReady {
		g.Go(func() error {
			if err := checkFelixHealth(ctx, felixReadinessEp, "readiness"); err != nil {
				return fmt.Errorf("calico/node is not ready: felix is not ready: %+v", err)
			}
			return nil
		})
	}

	if bird {
		g.Go(func() error {
			if err := checkBIRDReady("4", thresholdTime); err != nil {
				return fmt.Errorf("calico/node is not ready: BIRD is not ready: %+v", err)
			}
			return nil
		})
	}

	if bird6 {
		g.Go(func() error {
			if err := checkBIRDReady("6", thresholdTime); err != nil {
				return fmt.Errorf("calico/node is not ready: BIRD6 is not ready: %+v", err)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	 */
}