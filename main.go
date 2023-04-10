package main

import "github.com/Portfolio-Advanced-software/BingeBuster-MetadataService/messaging"

func main() {
	messaging.ProduceMessage("get description for movie 1", "metadata_movie")
	messaging.ConsumeMessage("metadata_movie")
}
