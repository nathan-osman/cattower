import { useEffect, useState } from 'react'
import BigButton from '../components/BigButton'
import VideoSvg from '../images/video.svg'

export default function Videos() {

  const [curVideo, setCurVideo] = useState<string | null>(null)
  const [videos, setVideos] = useState<string[]>([])

  useEffect(() => {
    fetch('/api/videos')
      .then(v => v.json())
      .then(v => setVideos(v))
  }, [])

  return (
    <>
      {curVideo !== null &&
        <video
          autoPlay
          className="absolute w-full h-full left-0 top-0 z-10"
          onEnded={() => setCurVideo(null)}
        >
          <source src={`/fs/videos/${curVideo}`} />
        </video>
      }
      <div className="flex">
        {videos.map((v, i) => (
          <BigButton
            key={i}
            icon={VideoSvg}
            text={v}
            onClick={() => setCurVideo(v)}
          />
        ))}
      </div>
    </>
  )
}
