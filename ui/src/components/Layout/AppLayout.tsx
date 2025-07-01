import type { FC } from "react"
import { Outlet } from 'react-router-dom'
import Topbar from './Topbar'
import Sidebar from './Sidebar'

const AppLayout: FC = () => {
  return (
    <div className="app">
      <Topbar />
      <div className="content-wrapper">
        <Sidebar />
        <main className="main-panel">
          <Outlet />
        </main>
      </div>
    </div>
  )
}

export default AppLayout
