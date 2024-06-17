import React, { useState, useEffect } from 'react';
import { useTable, useExpanded } from 'react-table';
import { Table, Button, FormControl } from 'react-bootstrap';
import '../styles/procesos.css';

function ProcessTable() {
  const [processList, setProcessList] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [pidSearchTerm, setPidSearchTerm] = useState('');
  const [systemInfo, setSystemInfo] = useState({});
  const [targetPid, setTargetPid] = useState('');
  const [searchMode, setSearchMode] = useState('name');

  const fetchProcesses = () => {
    fetch('http://192.168.122.15:3000/cpu')
      .then(response => response.json())
      .then(data => {
        if (data && data.processes) {
          setProcessList(data.processes);
          setSystemInfo(data.info || {});
        } else {
          setProcessList([]);
          setSystemInfo({});
        }
      })
      .catch(error => {
        console.error('Error fetching processes:', error);
        setProcessList([]);
        setSystemInfo({});
      });
  };

  useEffect(() => {
    fetchProcesses();
  }, []);

  const handleRefreshProcesses = () => {
    fetchProcesses();
  };

  const handleCreateProcess = () => {
    fetch('http://192.168.122.15:3000/cpu/iniProc/crear', {
      method: 'GET',
    }).then(() => {
      fetchProcesses();
      showAlert('Nuevo proceso creado');
    });
  };

  const handleKillProcess = () => {
    fetch(`http://192.168.122.15:3000/cpu/killProc?pid=${targetPid}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
    })
      .then(response => {
        if (response.ok) {
          return response.text();
        }
        throw new Error('Error al eliminar proceso');
      })
      .then(data => {
        showAlert(`Proceso con PID ${targetPid} ha terminado`);
        fetchProcesses();
      })
      .catch(error => {
        console.error('Error killing process:', error);
        showAlert('Error al eliminar proceso');
      });
  };

  const showAlert = (message) => {
    const alertContainer = document.createElement('div');
    alertContainer.className = 'alert-container';
    alertContainer.innerText = message;
    document.body.appendChild(alertContainer);

    setTimeout(() => {
      alertContainer.remove();
    }, 3000);
  };

  const getStateText = (state) => {
    switch (state) {
      case 1:
        return 'En ejecución';
      case 0:
        return 'Suspendido';
      case 128:
        return 'Detenido';
      case 1026:
        return 'Zombie';
      default:
        return state;
    }
  };

  const columns = React.useMemo(
    () => [
      {
        Header: 'PID',
        accessor: 'pid',
        Cell: ({ row }) => (
          <div>
            {row.original.pid}
            {row.original.child && row.original.child.length > 0 ? (
              <span {...row.getToggleRowExpandedProps()} style={{ marginLeft: '10px' }}>
                {row.isExpanded ? '▼' : '▶'}
              </span>
            ) : null}
          </div>
        ),
      },
      {
        Header: 'Nombre',
        accessor: 'name',
      },
      {
        Header: '#Estado',
        accessor: 'state',
        id: 'stateNumber',
      },
      {
        Header: 'Estado',
        accessor: 'state',
        id: 'stateText',
        Cell: ({ value }) => getStateText(value),
      },
    ],
    []
  );

  const data = React.useMemo(
    () =>
      processList.map(process => ({
        ...process,
        subRows: process.child || [],
      })),
    [processList]
  );

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow,
  } = useTable(
    {
      columns,
      data,
    },
    useExpanded
  );

  const handleSearchChange = event => {
    setSearchTerm(event.target.value);
  };

  const handlePidSearchChange = event => {
    setPidSearchTerm(event.target.value);
  };

  const toggleSearchMode = () => {
    setSearchMode(prevMode => (prevMode === 'name' ? 'pid' : 'name'));
  };

  const filteredRows = rows.filter(row => {
    if (searchMode === 'name') {
      return row.original.name.toLowerCase().includes(searchTerm.toLowerCase());
    }
    return row.original.pid.toString().includes(pidSearchTerm);
  });

  return (
    <div className="process-table-container">
      <div className="form-container">
        <div className="title-container">
          <h3>Ingrese el PID</h3>
        </div>
        <FormControl
          placeholder="xxxxxxxxxx"
          value={targetPid}
          onChange={(e) => setTargetPid(e.target.value)}
        />
      </div>
      <div className="action-buttons">
        <Button variant="danger" onClick={handleKillProcess}>
          <i className="bi bi-file-earmark-x"></i> Matar Proceso
        </Button>
        <Button variant="success" onClick={handleCreateProcess}>
          <i className="bi bi-file-earmark-plus"></i> Crear Proceso
        </Button>
      </div>
      <div className="table-container">
        <Table striped bordered hover {...getTableProps()}>
          <thead>
            {headerGroups.map(headerGroup => (
              <tr {...headerGroup.getHeaderGroupProps()}>
                {headerGroup.headers.map(column => (
                  <th {...column.getHeaderProps()}>{column.render('Header')}</th>
                ))}
              </tr>
            ))}
          </thead>
          <tbody {...getTableBodyProps()}>
            {filteredRows.map(row => {
              prepareRow(row);
              return (
                <React.Fragment key={row.id}>
                  <tr {...row.getRowProps()}>
                    {row.cells.map(cell => (
                      <td {...cell.getCellProps()}>{cell.render('Cell')}</td>
                    ))}
                  </tr>
                  {row.isExpanded && row.original.child.length > 0 && (
                    <tr>
                      <td colSpan={columns.length}>
                        <Table striped bordered hover>
                          <tbody>
                            {row.original.child.map(child => (
                              <tr key={child.pid}>
                                <td>{child.pid}</td>
                                <td>{child.name}</td>
                                <td>{child.state}</td>
                                <td>{getStateText(child.state)}</td>
                              </tr>
                            ))}
                          </tbody>
                        </Table>
                      </td>
                    </tr>
                  )}
                </React.Fragment>
              );
            })}
          </tbody>
        </Table>
      </div>
    </div>
  );
}

export default ProcessTable;
