package consts

import "time"

const Support = "email@example.com"
const PublicPort = 38080
const Domain = "localhost"

const CleanupFrequency = 10 * time.Minute // frequency of periodic room cleanup
const MaximumRoomAge = 12 * time.Hour     // duration after a room is considered inactive
