import './styles/App.css'
import Head from "./components/Encabezado"
import RealTimeCharts from './components/Graficas'
import ProcessTable from './components/procesos'


function App() {
  
  let component
  switch (window.location.pathname) {
    case "/":
      component = <RealTimeCharts />
      break;
    // case "/cpuyram":
    //   component = <RealTimeCharts />
      // break;
    case "/cpu":
      component = <ProcessTable />
      break;
    default:
  }

  return (
    <>
    <Head />
    {component}
    </>
  )
  
}

export default App
