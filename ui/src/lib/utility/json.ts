export const format = (data: string): string => {
  try {
    return JSON.stringify(JSON.parse(data), null, '  ')
  } catch (e) {}

  return data
}
