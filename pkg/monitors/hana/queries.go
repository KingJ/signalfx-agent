package hana

import "github.com/signalfx/signalfx-agent/pkg/monitors/sql"

// Queries that get metrics about the entire server instance and do not need to
// be run on a per-database basis.
var defaultServerQueries = []sql.Query{
	{
		Query: `SELECT host AS hana_host, usage_type, used_size FROM m_disk_usage WHERE used_size >= 0;`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.disk.used_size",
				ValueColumn:      "used_size",
				DimensionColumns: []string{"hana_host", "usage_type"},
			},
		},
	},
	{
		Query: `SELECT host AS hana_host, SUM(total_device_size) AS total_size FROM (SELECT device_id, host, MAX(total_device_size) AS total_device_size FROM m_disks GROUP BY device_id, host) GROUP BY host;`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.disk.total_size",
				ValueColumn:      "total_size",
				DimensionColumns: []string{"hana_host"},
			},
		},
	},
	{
		Query: `SELECT host AS hana_host, service_name, process_cpu, open_file_count FROM m_service_statistics;`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.service.cpu.utilization",
				ValueColumn:      "process_cpu",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.file.open",
				ValueColumn:      "open_file_count",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
		},
	},
	{
		Query: `SELECT host AS hana_host, free_physical_memory, used_physical_memory, free_swap_space, used_swap_space, allocation_limit, instance_total_memory_used_size, instance_total_memory_allocated_size, instance_code_size, instance_shared_memory_allocated_size, open_file_count FROM m_host_resource_utilization;`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.host.memory.physical.free",
				ValueColumn:      "free_physical_memory",
				DimensionColumns: []string{"hana_host"},
			},
			{
				MetricName:       "sap.hana.host.memory.physical.used",
				ValueColumn:      "used_physical_memory",
				DimensionColumns: []string{"hana_host"},
			},
			{
				MetricName:       "sap.hana.host.memory.swap.free",
				ValueColumn:      "free_swap_space",
				DimensionColumns: []string{"hana_host"},
			},
			{
				MetricName:       "sap.hana.host.memory.swap.used",
				ValueColumn:      "used_swap_space",
				DimensionColumns: []string{"hana_host"},
			},
			{
				MetricName:       "sap.hana.host.memory.allocation.limit",
				ValueColumn:      "allocation_limit",
				DimensionColumns: []string{"hana_host"},
			},
			{
				MetricName:       "sap.hana.host.memory.poll.used",
				ValueColumn:      "instance_total_memory_used_size",
				DimensionColumns: []string{"hana_host"},
			},
			{
				MetricName:       "sap.hana.host.memory.poll.total",
				ValueColumn:      "instance_total_memory_allocated_size",
				DimensionColumns: []string{"hana_host"},
			},
			{
				MetricName:       "sap.hana.host.memory.code",
				ValueColumn:      "instance_code_size",
				DimensionColumns: []string{"hana_host"},
			},
			{
				MetricName:       "sap.hana.host.memory.shared",
				ValueColumn:      "instance_shared_memory_allocated_size",
				DimensionColumns: []string{"hana_host"},
			},
			{
				MetricName:       "sap.hana.host.file.open",
				ValueColumn:      "open_file_count",
				DimensionColumns: []string{"hana_host"},
			},
		},
	},
	{
		Query: `SELECT host AS hana_host, service_name, logical_memory_size, physical_memory_size, code_size, stack_size, heap_memory_allocated_size, heap_memory_used_size, shared_memory_allocated_size, shared_memory_used_size, allocation_limit, effective_allocation_limit, total_memory_used_size FROM m_service_memory;`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.service.memory.logical",
				ValueColumn:      "logical_memory_size",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.physical",
				ValueColumn:      "physical_memory_size",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.code",
				ValueColumn:      "code_size",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.stack",
				ValueColumn:      "stack_size",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.heap.allocated",
				ValueColumn:      "heap_memory_allocated_size",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.heap.used",
				ValueColumn:      "heap_memory_used_size",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.shared.allocated",
				ValueColumn:      "shared_memory_allocated_size",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.shared.used",
				ValueColumn:      "shared_memory_used_size",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.pool.limit",
				ValueColumn:      "allocation_limit",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.pool.limit.effective",
				ValueColumn:      "effective_allocation_limit",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
			{
				MetricName:       "sap.hana.service.memory.pool.used",
				ValueColumn:      "total_memory_used_size",
				DimensionColumns: []string{"hana_host", "service_name"},
			},
		},
	},
	{
		Query: `SELECT services.host AS hana_host, services.service_name AS service_name, memory.component AS component_name, memory.used_memory_size AS used_memory_size FROM m_service_component_memory AS memory JOIN m_services AS services ON memory.host = services.host AND memory.port = services.port;`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.service.component.used",
				ValueColumn:      "used_memory_size",
				DimensionColumns: []string{"hana_host", "service_name", "component_name"},
			},
		},
	},
	{
		Query: `SELECT host AS hana_host, statement_hash, SUM(recompile_count) AS recompile_count, SUM(execution_count) AS execution_count, TO_DOUBLE(AVG(avg_execution_time)) AS avg_execution_time, MIN(min_execution_time) AS min_execution_time, MAX(max_execution_time) AS max_execution_time, SUM(total_execution_time) AS total_execution_time FROM m_active_statements GROUP BY host, statement_hash;`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.statement.active.recompile.count",
				ValueColumn:      "recompile_count",
				DimensionColumns: []string{"hana_host", "statement_hash"},
			},
			{
				MetricName:       "sap.hana.statement.active.execution.count",
				ValueColumn:      "execution_count",
				DimensionColumns: []string{"hana_host", "statement_hash"},
			},
			{
				MetricName:       "sap.hana.statement.active.execution.mean",
				ValueColumn:      "avg_execution_time",
				DimensionColumns: []string{"hana_host", "statement_hash"},
			},
			{
				MetricName:       "sap.hana.statement.active.execution.sum",
				ValueColumn:      "total_execution_time",
				DimensionColumns: []string{"hana_host", "statement_hash"},
			},
			{
				MetricName:       "sap.hana.statement.active.execution.min",
				ValueColumn:      "min_execution_time",
				DimensionColumns: []string{"hana_host", "statement_hash"},
			},
			{
				MetricName:       "sap.hana.statement.active.execution.max",
				ValueColumn:      "max_execution_time",
				DimensionColumns: []string{"hana_host", "statement_hash"},
			},
		},
	},
	{
		Query: `SELECT host AS hana_host, statement_hash, db_user, schema_name, app_user, operation, duration_microsec, records FROM m_expensive_statements;`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.statement.expensive.duration",
				ValueColumn:      "duration_microsec",
				DimensionColumns: []string{"hana_host", "statement_hash", "db_user", "schema_name", "app_user", "operation"},
			},
			{
				MetricName:       "sap.hana.statement.expensive.records",
				ValueColumn:      "records",
				DimensionColumns: []string{"hana_host", "statement_hash", "db_user", "schema_name", "app_user", "operation"},
			},
		},
	},
	{
		Query: `SELECT host AS hana_host, statement_hash, db_user, schema_name, app_user, operation, COUNT(*) AS errors FROM m_expensive_statements WHERE error_code <> 0 GROUP BY host, statement_hash, db_user, schema_name, app_user, operation;`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.statement.expensive.errors",
				ValueColumn:      "errors",
				DimensionColumns: []string{"hana_host", "statement_hash", "db_user", "schema_name", "app_user", "operation"},
			},
		},
	},
	{
		Query: `SELECT host AS hana_host, connection_status, user_name, COUNT(*) AS count, SUM(memory_size_per_connection) AS memory_size, SUM(fetched_record_count) AS fetched_record_count, SUM(affected_record_count) AS affected_record_count, SUM(sent_message_size) AS sent_message_size, SUM(sent_message_count) AS sent_message_count, SUM(received_message_size) AS received_message_size, SUM(received_message_count) AS received_message_count FROM m_connections GROUP BY host, connection_status, user_name HAVING connection_status != '';`,
		Metrics: []sql.Metric{
			{
				MetricName:       "sap.hana.connection.count",
				ValueColumn:      "count",
				DimensionColumns: []string{"hana_host", "connection_status", "user_name"},
			},
			{
				MetricName:       "sap.hana.connection.memory.allocated",
				ValueColumn:      "memory_size",
				DimensionColumns: []string{"hana_host", "connection_status", "user_name"},
			},
			{
				MetricName:       "sap.hana.connection.record.fetched",
				ValueColumn:      "fetched_record_count",
				IsCumulative:     true,
				DimensionColumns: []string{"hana_host", "connection_status", "user_name"},
			},
			{
				MetricName:       "sap.hana.connection.record.affected",
				ValueColumn:      "affected_record_count",
				IsCumulative:     true,
				DimensionColumns: []string{"hana_host", "connection_status", "user_name"},
			},
			{
				MetricName:       "sap.hana.connection.message.sent.size",
				ValueColumn:      "sent_message_size",
				IsCumulative:     true,
				DimensionColumns: []string{"hana_host", "connection_status", "user_name"},
			},
			{
				MetricName:       "sap.hana.connection.message.sent.count",
				ValueColumn:      "sent_message_count",
				IsCumulative:     true,
				DimensionColumns: []string{"hana_host", "connection_status", "user_name"},
			},
			{
				MetricName:       "sap.hana.connection.message.received.size",
				ValueColumn:      "received_message_size",
				IsCumulative:     true,
				DimensionColumns: []string{"hana_host", "connection_status", "user_name"},
			},
			{
				MetricName:       "sap.hana.connection.message.received.count",
				ValueColumn:      "received_message_count",
				IsCumulative:     true,
				DimensionColumns: []string{"hana_host", "connection_status", "user_name"},
			},
		},
	},
}
