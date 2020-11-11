// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"focus/app/dao/internal"
)

// zanCaiDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type zanCaiDao struct {
	*internal.ZanCaiDao
}

var (
	// ZanCai is globally public accessible object for table {TplTableName} operations.
	ZanCai = &zanCaiDao{
		internal.ZanCai,
	}
)

// Fill with you ideas below.