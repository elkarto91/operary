import { BrowserRouter, Route, Routes } from 'react-router-dom'
import AppLayout from './components/Layout/AppLayout'
import Dashboard from './pages/Dashboard'
import CorePad from './pages/CorePad'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<AppLayout />}> 
          <Route index element={<Dashboard />} />
          <Route path="corepad" element={<CorePad />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
