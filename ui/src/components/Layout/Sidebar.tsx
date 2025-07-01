import type { FC } from "react"
import { NavLink } from 'react-router-dom'
import './sidebar.css'

const Sidebar: FC = () => {
  return (
    <aside className="sidebar">
      <nav>
        <NavLink to="/" end>Dashboard</NavLink>
        <NavLink to="/corepad">Notes</NavLink>
      </nav>
    </aside>
  )
}

export default Sidebar
