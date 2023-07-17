// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package hrsMdl

import (
	"database/sql"
)

type HrDepartment struct {
	DeptID           sql.NullInt32          `db:"dept_id" json:"deptId"`
	DeptName         sql.NullString `db:"dept_name" json:"deptName"`
	DeptModifiedDate sql.NullTime   `db:"dept_modified_date" json:"deptModifiedDate"`
}

type HrEmployee struct {
	EmpEntityID       sql.NullInt32          `db:"emp_entity_id" json:"empEntityId"`
	EmpEmpNumber      sql.NullString `db:"emp_emp_number" json:"empEmpNumber"`
	EmpNationalID     sql.NullString `db:"emp_national_id" json:"empNationalId"`
	EmpBirthDate      sql.NullTime   `db:"emp_birth_date" json:"empBirthDate"`
	EmpMaritalStatus  sql.NullString `db:"emp_marital_status" json:"empMaritalStatus"`
	EmpGender         sql.NullString `db:"emp_gender" json:"empGender"`
	EmpHireDate       sql.NullTime   `db:"emp_hire_date" json:"empHireDate"`
	EmpSalariedFlag   sql.NullString `db:"emp_salaried_flag" json:"empSalariedFlag"`
	EmpVacationHours  sql.NullInt16  `db:"emp_vacation_hours" json:"empVacationHours"`
	EmpSickleaveHours sql.NullInt16  `db:"emp_sickleave_hours" json:"empSickleaveHours"`
	EmpCurrentFlag    sql.NullInt16  `db:"emp_current_flag" json:"empCurrentFlag"`
	EmpModifiedDate   sql.NullTime   `db:"emp_modified_date" json:"empModifiedDate"`
	EmpType           sql.NullString `db:"emp_type" json:"empType"`
	EmpJoroID         sql.NullInt32  `db:"emp_joro_id" json:"empJoroId"`
	EmpEmpEntityID    sql.NullInt32  `db:"emp_emp_entity_id" json:"empEmpEntityId"`
}

type HrEmployeeClientContract struct {
	EccoID             sql.NullInt32          `db:"ecco_id" json:"eccoId"`
	EccoEntityID       sql.NullInt32          `db:"ecco_entity_id" json:"eccoEntityId"`
	EccoContractNo     sql.NullString `db:"ecco_contract_no" json:"eccoContractNo"`
	EccoContractDate   sql.NullTime   `db:"ecco_contract_date" json:"eccoContractDate"`
	EccoStartDate      sql.NullTime   `db:"ecco_start_date" json:"eccoStartDate"`
	EccoEndDate        sql.NullTime   `db:"ecco_end_date" json:"eccoEndDate"`
	EccoNotes          sql.NullString `db:"ecco_notes" json:"eccoNotes"`
	EccoModifiedDate   sql.NullTime   `db:"ecco_modified_date" json:"eccoModifiedDate"`
	EccoMediaLink      sql.NullString `db:"ecco_media_link" json:"eccoMediaLink"`
	EccoJotyID         sql.NullInt32  `db:"ecco_joty_id" json:"eccoJotyId"`
	EccoAccountManager sql.NullInt32  `db:"ecco_account_manager" json:"eccoAccountManager"`
	EccoClitID         sql.NullInt32  `db:"ecco_clit_id" json:"eccoClitId"`
	EccoStatus         sql.NullString `db:"ecco_status" json:"eccoStatus"`
}

type HrEmployeeDepartmentHistory struct {
	EdhiID           sql.NullInt32         `db:"edhi_id" json:"edhiId"`
	EdhiEntityID     sql.NullInt32         `db:"edhi_entity_id" json:"edhiEntityId"`
	EdhiStartDate    sql.NullTime  `db:"edhi_start_date" json:"edhiStartDate"`
	EdhiEndDate      sql.NullTime  `db:"edhi_end_date" json:"edhiEndDate"`
	EdhiModifiedDate sql.NullTime  `db:"edhi_modified_date" json:"edhiModifiedDate"`
	EdhiDeptID       sql.NullInt32 `db:"edhi_dept_id" json:"edhiDeptId"`
}

type HrEmployeePayHistory struct {
	EphiEntityID       sql.NullInt32         `db:"ephi_entity_id" json:"ephiEntityId"`
	EphiRateChangeDate sql.NullTime     `db:"ephi_rate_change_date" json:"ephiRateChangeDate"`
	EphiRateSalary     sql.NullInt32 `db:"ephi_rate_salary" json:"ephiRateSalary"`
	EphiPayFrequence   sql.NullInt16 `db:"ephi_pay_frequence" json:"ephiPayFrequence"`
	EphiModifiedDate   sql.NullTime  `db:"ephi_modified_date" json:"ephiModifiedDate"`
}
