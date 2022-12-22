package impl

const (
	InsertResourceSQL = `
  INSERT INTO resource (
    id,
    vendor,
    region,
    create_at,
    expire_at,
    type,
    name,
    description,
    status,
    update_at,
    sync_at,
    accout,
    public_ip,
    private_ip
  ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?);
`
	InsertDescribeSQL = `
    INSERT INTO host (
    resource_id,
    cpu,
    memory,
    gpu_amount,
    gpu_spec,
    os_type,
    os_name,
    serial_number
  ) VALUES (?,?,?,?,?,?,?,?);
`

	QueryHostSQL = `
  SELECT
    r.*, h.cpu, h.memory, h.gpu_amount, h.gpu_spec, h.os_type, h.os_name, h.serial_number
  FROM
    resource AS r
    LEFT JOIN host as h ON r.id = h.resource_id
  `
)
