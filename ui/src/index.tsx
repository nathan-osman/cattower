import * as React from 'react'
import * as ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from 'react-router-dom'
import App from './components/App'
import Index from './routes'
import Leds from './routes/leds'
import Sensors from './routes/sensors'
import Videos from './routes/videos'
import './index.css'

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    children: [
      {
        path: '/',
        element: <Index />,
      },
      {
        path: '/leds',
        element: <Leds />,
      },
      {
        path: '/sensors',
        element: <Sensors />,
      },
      {
        path: '/videos',
        element: <Videos />,
      },
    ],
  },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)
