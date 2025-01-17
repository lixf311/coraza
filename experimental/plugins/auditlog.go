// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package plugins

import (
	"github.com/lixf311/coraza/v3/experimental/plugins/plugintypes"
	"github.com/lixf311/coraza/v3/internal/auditlog"
)

// RegisterAuditLogWriter registers a new audit log writer.
func RegisterAuditLogWriter(name string, writerFactory func() plugintypes.AuditLogWriter) {
	auditlog.RegisterWriter(name, writerFactory)
}

// RegisterAuditLogFormatter registers a new audit log formatter.
func RegisterAuditLogFormatter(name string, format plugintypes.AuditLogFormatter) {
	auditlog.RegisterFormatter(name, format)
}
