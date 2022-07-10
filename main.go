package main

import (
	scheduleactive "github.com/chaos-mesh/chaos-mesh/controllers/schedule/active"
	"github.com/chaos-mesh/chaos-mesh/controllers/schedule/utils"
	"github.com/chaos-mesh/chaos-mesh/controllers/types"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/recorder"
	"github.com/chaos-mesh/chaos-mesh/pkg/metrics"
	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	manager, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{})
	if err != nil {
		os.Exit(1)
	}

	clientactive := scheduleactive.Reconciler{Client: manager.GetClient()}
	//clientcron := schedulecron.Reconciler{Client: manager.GetClient()}
	//clientgc := schedulegc.Reconciler{Client: manager.GetClient()}
	//clientpause := schedulepause.Reconciler{Client: manager.GetClient()}

	ChaosControllerManagerMetricsCollector := metrics.NewChaosControllerManagerMetricsCollector(manager, prometheus.NewRegistry(), logr.Logger{})
	recorderBuilder := recorder.NewRecorderBuilder(clientactive, logr.Logger{}, runtime.NewScheme(), ChaosControllerManagerMetricsCollector)
	scheduleactive.Bootstrap(manager, clientactive, logr.Logger{}, scheduleactive.Objs{Objs: []types.Object{}}, runtime.NewScheme(), utils.NewActiveLister(clientactive, logr.Logger{}), recorderBuilder)

	//schedulecron.Bootstrap(manager, clientcron)
	//schedulegc.Bootstrap(manager, clientgc)
	//schedulepause.Bootstrap(manager, clientpause)

}
