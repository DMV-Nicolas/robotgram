export const getTimeElapsed = (time: string) => {
  // TODO: get time elapsed with hours, minutes and seconds
  const targetDate = new Date(time)
  const actualDate = new Date()

  const daysDiff = actualDate.getTime() - targetDate.getTime()
  const days = Math.round(daysDiff / (1000 * 60 * 60 * 24))
  const hours = Math.round(daysDiff / (1000 * 60 * 60))
  const minutes = Math.round(daysDiff / (1000 * 60))
  const seconds = Math.round(daysDiff / (1000))

  if (days !== 0) {
    return `${days} d`
  } else if (hours !== 0) {
    return `${hours} h`
  } else if (minutes !== 0) {
    return `${minutes} m`
  } else if (seconds !== 0) {
    return `${seconds} s`
  } else {
    return 'now'
  }
}
