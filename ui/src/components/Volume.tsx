import { ChangeEvent, FocusEvent, useEffect, useState } from 'react'
import VolumeSvg from '../images/volume.svg'
import VolumeDownSvg from '../images/volume-down.svg'
import VolumeMutedSvg from '../images/volume-muted.svg'
import VolumeUpSvg from '../images/volume-up.svg'

type Volume = {
  volume: number
}

export default function Volume() {

  const [volume, setVolume] = useState(0)
  const [overlay, setOverlay] = useState(false)

  useEffect(() => {
    fetch('/api/system/volume')
      .then(v => v.json())
      .then(v => setVolume(v.volume))
  }, [])

  let timeoutId: NodeJS.Timeout

  function handleFocus() {
    clearTimeout(timeoutId)
  }

  function handleBlur(e: FocusEvent) {
    timeoutId = setTimeout(() => {
      setOverlay(false)
    }, 0)
  }

  function updateVolume(v: number) {
    setVolume(v)
    fetch('/api/system/volume/set', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
      },
      body: JSON.stringify({
        volume: v,
      }),
    })
  }

  function handleChange(e: ChangeEvent<HTMLInputElement>) {
    updateVolume(Number(e.target.value))
  }

  return (
    <div
      className="relative"
      onFocus={handleFocus}
      onBlur={handleBlur}
    >
      <img
        src={volume ? VolumeSvg : VolumeMutedSvg}
        className="w-12"
        onClick={() => setOverlay(!overlay)}
      />
      {overlay &&
        <div className="border border-foreground bg-background absolute w-[480px] right-0 p-2 flex gap-x-4">
          <button
            type="button"
            onClick={() => updateVolume(Math.max(0, volume - 2))}
          >
            <img src={VolumeDownSvg} className="w-12" />
          </button>
          <input
            type="range"
            min="0"
            max="100"
            className="grow"
            value={volume}
            onChange={handleChange}
          />
          <button
            onClick={() => updateVolume(Math.min(100, volume + 2))}
          >
            <img src={VolumeUpSvg} className="w-12" />
          </button>
        </div>
      }
    </div>
  )
}
