package util

// ValidChannelName ...
func ValidChannelName(channel string) bool {
	return true
}

// NormalizeChannelName ...
func NormalizeChannelName(channel string) string {
	return channel
}

/*
// TODO: port from node-streamhut

const reservedWords = require('./reserved_words.json')

function isReservedKeyword(s) {
  s = s.toLowerCase().trim()
  return reservedWords.indexOf(s) > -1
}

function normalizeChannel(channel) {
  return channel.replace(/[^a-zA-Z0-9-]+/, '').toLowerCase().trim()
}

function isValidChannelName(channel) {
  if (isReservedKeyword(channel)) {
    return false
  }

  return /^[a-zA-Z0-9-]+$/.test(channel)
}

module.exports = {
  isReservedKeyword,
  normalizeChannel,
  isValidChannelName
}
*/
