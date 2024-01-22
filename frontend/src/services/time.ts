export const getTimeElapsed = (time: string) => {
  // TODO: get time elapsed with hours, minutes and seconds
  const targetDate = new Date(time)
  const actualDate = new Date()

  const daysDiff = actualDate.getTime() - targetDate.getTime()
  const days = Math.round(daysDiff / (1000 * 60 * 60 * 24))
  return `${days} d`
}
