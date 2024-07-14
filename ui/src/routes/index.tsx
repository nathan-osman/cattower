import { useNavigate } from 'react-router-dom'
import BigButton from '../components/BigButton'
import SensorOverview from '../components/SensorOverview'
import LampSvg from '../images/lamp.svg'
import SensorsSvg from '../images/sensors.svg'
import VideosSvg from '../images/videos.svg'

export default function Index() {

  const navigate = useNavigate()

  return (
    <div className="flex gap-x-2 h-full">
      <div className="grow">
        <SensorOverview />
      </div>
      <div className="flex gap-x-2 items-start">
        <BigButton
          icon={LampSvg}
          text="LED Control"
          onClick={() => navigate('/leds')}
        />
        <BigButton
          icon={SensorsSvg}
          text="Sensors"
          onClick={() => navigate('/sensors')}
        />
        <BigButton
          icon={VideosSvg}
          text="Videos"
          onClick={() => navigate('/videos')}
        />
      </div>
    </div>
  )
}
