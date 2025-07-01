package handlers

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	startTime     = time.Now()
	totalRequests uint64
	shiftsStarted uint64
	shiftsClosed  uint64
	tasksCreated  uint64
)
var (
	AuditSyncTotalSubmitted = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "auditsync_total_submitted",
			Help: "Total number of audit entries submitted",
		},
	)
	AuditSyncSubmissionFailures = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "auditsync_submission_failures_total",
			Help: "Number of failed audit submissions",
		},
	)
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&totalRequests, 1)
		next.ServeHTTP(w, r)
	})
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).Seconds()
	count := atomic.LoadUint64(&totalRequests)
	started := atomic.LoadUint64(&shiftsStarted)
	closed := atomic.LoadUint64(&shiftsClosed)
	created := atomic.LoadUint64(&tasksCreated)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "operary_uptime_seconds %.2f\n", uptime)
	fmt.Fprintf(w, "operary_total_requests %d\n", count)
	fmt.Fprintf(w, "operary_shifts_started %d\n", started)
	fmt.Fprintf(w, "operary_shifts_closed %d\n", closed)
	fmt.Fprintf(w, "operary_tasks_created %d\n", created)
}

var (
	CorePadNotesCreated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "corepad_notes_created_total",
			Help: "Total number of notes created in CorePad",
		},
	)
	CorePadNotesFailed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "corepad_notes_failed_total",
			Help: "Number of failed note creations",
		},
	)
)

var OpsMirrorWarnings = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "opsmirror_status_warnings_total",
		Help: "Number of systems in WARN state",
	},
)

var (
	PermitGridTotalSubmitted = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "permitgrid_total_submitted",
			Help: "Total number of permit requests submitted",
		},
	)

	PermitGridApprovals = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "permitgrid_total_approved",
			Help: "Total number of permits approved",
		},
	)

	PermitGridRejections = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "permitgrid_total_rejected",
			Help: "Total number of permits rejected",
		},
	)
)
var (
	EquipTrustCheckouts = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "equiptrust_checkouts_total",
			Help: "Total number of equipment check-outs",
		},
	)

	EquipTrustReturns = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "equiptrust_returns_total",
			Help: "Total number of equipment returns",
		},
	)

	TrainOpsSamplesTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "trainops_samples_total",
			Help: "Total training samples ingested",
		},
	)
)

func init() {
	prometheus.MustRegister(CorePadNotesCreated)
	prometheus.MustRegister(CorePadNotesFailed)
	prometheus.MustRegister(OpsMirrorWarnings)
	prometheus.MustRegister(AuditSyncTotalSubmitted)
	prometheus.MustRegister(AuditSyncSubmissionFailures)
	prometheus.MustRegister(PermitGridTotalSubmitted)
	prometheus.MustRegister(PermitGridApprovals)
	prometheus.MustRegister(PermitGridRejections)
	prometheus.MustRegister(EquipTrustCheckouts)
	prometheus.MustRegister(EquipTrustReturns)
	prometheus.MustRegister(TrainOpsSamplesTotal)
}
