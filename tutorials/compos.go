package main
// Board represents a surface we can work on.
 type Board struct {
     NailsNeeded int
     NailsDriven int
 }

// NailDriver represents behavior to drive nails into a board.
 type NailDriver interface {
     DriveNail(nailSupply *int, b *Board)
 }

 // NailPuller represents behavior to remove nails into a board.
 type NailPuller interface {
     PullNail(nailSupply *int, b *Board)
 }

 // NailDrivePuller represents behavior to drive/remove nails into a board.
 type NailDrivePuller interface {
     NailDriver
     NailPuller
 }

// Mallet is a tool that pounds in nails.
 type Mallet struct{}

 // DriveNail pounds a nail into the specified board.
 func (Mallet) DriveNail(nailSupply *int, b *Board) {
     // Take a nail out of the supply.
     *nailSupply–

     // Pound a nail into the board.
     b.NailsDriven++

     fmt.Println("Mallet: pounded nail into the board.")
 }
