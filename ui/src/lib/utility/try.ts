export const tryunwrap = <T>(fn: () => void, ok: T, fail: T) => {
  try {
    fn()
  } catch (e) {
    return fail
  }
  return ok
}
