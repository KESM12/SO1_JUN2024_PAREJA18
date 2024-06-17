import React, { useEffect, useState } from 'react';
import { Pie } from 'react-chartjs-2';
import 'chart.js/auto'; // Importación automática de los componentes de Chart.js
import '../styles/Estilo.css'; // Asegúrate de crear este archivo para los estilos

// Aqui se grafica el uso de CPU y RAM en tiempo real

function RealTimeCharts() {
  const [cpuUsage, setCpuUsage] = useState(null);
  const [ramUsage, setRamUsage] = useState(null);
  const url = "/back"; // Cambiar por la URL de tu API

  useEffect(() => {
    const fetchUsageData = () => {
      // Fetch data from the API
      fetch(url + 'cpuyram') // Reemplaza con tu endpoint real // Quite la ip y coloque localhost para validar si jalaba o no.
        .then(response => response.json())
        .then(data => {
          setCpuUsage(data.cpu_percentage);
          setRamUsage(data.ram_percentage);
          console.log('Datos recibidos:', data); // Imprimir en la consola
        })
        .catch(error => console.error('Error fetching data:', error));
    };

    fetchUsageData(); // Realiza la primera llamada

    const interval = setInterval(() => {
      fetchUsageData(); // Realiza una llamada cada 2 segundos
    }, 2000);

    return () => clearInterval(interval); // Limpia el intervalo al desmontar el componente
  }, []);

  const generatePieData = (label, percentage) => {
    // Convertir el porcentaje a un número flotante
    const percentageValue = parseFloat(percentage);
    // Si el porcentaje es 0, mostrar la gráfica como 100% libre
    const data = percentageValue === 0 ? [0, 100] : [percentageValue, 100 - percentageValue];
  
    return {
      labels: [label, 'Libre'],
      datasets: [
        {
          data: data,
          backgroundColor: percentageValue === 0 ? ['#BAFF39', '#6E6E6E'] : ['#BAFF39', '#6E6E6E'], // Ajustar colores según el caso
        },
      ],
    };
  };

  return (
    <div className="statistics-container">
      <div className="title-container">
        <h1>SO1 - JUN 2024</h1>
      </div>
      <div className="charts-container">
        {cpuUsage !== null && (
          <div className="chart">
            <h2>{cpuUsage} %CPU</h2>
            <Pie data={generatePieData('CPU', cpuUsage)} />
          </div>
        )}
        {ramUsage !== null && (
          <div className="chart">
            <h2>{ramUsage} %RAM</h2>
            <Pie data={generatePieData('RAM', ramUsage)} />
          </div>
        )}
      </div>
    </div>
  );
}

export default RealTimeCharts;
