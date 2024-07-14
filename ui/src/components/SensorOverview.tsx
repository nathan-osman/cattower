import { useEffect, useState } from 'react'

type Measurement = {
  name: string
  value: number
}

export default function SensorOverview() {

  const [values, setValues] = useState<Measurement[]>([])

  useEffect(() => {
    const updateValues = () => {
      fetch('/api/sensors/overview')
        .then(v => v.json())
        .then(v => setValues(v))
    }
    updateValues()
    const id = setInterval(updateValues, 300 * 1000)
    return () => {
      clearInterval(id)
    }
  }, [])

  return (
    <table className="border-separate border-spacing-x-2">
      {
        values.map(v => (
          <tr>
            <th className="text-right">
              {v.name}:
            </th>
            <td>{v.value.toFixed(1)} &deg;C</td>
          </tr>
        ))
      }
    </table>
  )
}
