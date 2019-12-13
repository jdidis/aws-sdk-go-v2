// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type StartGameSessionPlacementInput struct {
	_ struct{} `type:"structure"`

	// Set of information on each player to create a player session for.
	DesiredPlayerSessions []DesiredPlayerSession `type:"list"`

	// Set of custom properties for a game session, formatted as key:value pairs.
	// These properties are passed to a game server process in the GameSession object
	// with a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	GameProperties []GameProperty `type:"list"`

	// Set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process in the GameSession object with
	// a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	GameSessionData *string `min:"1" type:"string"`

	// Descriptive label that is associated with a game session. Session names do
	// not need to be unique.
	GameSessionName *string `min:"1" type:"string"`

	// Name of the queue to use to place the new game session.
	//
	// GameSessionQueueName is a required field
	GameSessionQueueName *string `min:"1" type:"string" required:"true"`

	// Maximum number of players that can be connected simultaneously to the game
	// session.
	//
	// MaximumPlayerSessionCount is a required field
	MaximumPlayerSessionCount *int64 `type:"integer" required:"true"`

	// Unique identifier to assign to the new game session placement. This value
	// is developer-defined. The value must be unique across all regions and cannot
	// be reused unless you are resubmitting a canceled or timed-out placement request.
	//
	// PlacementId is a required field
	PlacementId *string `min:"1" type:"string" required:"true"`

	// Set of values, expressed in milliseconds, indicating the amount of latency
	// that a player experiences when connected to AWS regions. This information
	// is used to try to place the new game session where it can offer the best
	// possible gameplay experience for the players.
	PlayerLatencies []PlayerLatency `type:"list"`
}

// String returns the string representation
func (s StartGameSessionPlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartGameSessionPlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartGameSessionPlacementInput"}
	if s.GameSessionData != nil && len(*s.GameSessionData) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionData", 1))
	}
	if s.GameSessionName != nil && len(*s.GameSessionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionName", 1))
	}

	if s.GameSessionQueueName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameSessionQueueName"))
	}
	if s.GameSessionQueueName != nil && len(*s.GameSessionQueueName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionQueueName", 1))
	}

	if s.MaximumPlayerSessionCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaximumPlayerSessionCount"))
	}

	if s.PlacementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlacementId"))
	}
	if s.PlacementId != nil && len(*s.PlacementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementId", 1))
	}
	if s.DesiredPlayerSessions != nil {
		for i, v := range s.DesiredPlayerSessions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DesiredPlayerSessions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.GameProperties != nil {
		for i, v := range s.GameProperties {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GameProperties", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.PlayerLatencies != nil {
		for i, v := range s.PlayerLatencies {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PlayerLatencies", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type StartGameSessionPlacementOutput struct {
	_ struct{} `type:"structure"`

	// Object that describes the newly created game session placement. This object
	// includes all the information provided in the request, as well as start/end
	// time stamps and placement status.
	GameSessionPlacement *GameSessionPlacement `type:"structure"`
}

// String returns the string representation
func (s StartGameSessionPlacementOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartGameSessionPlacement = "StartGameSessionPlacement"

// StartGameSessionPlacementRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Places a request for a new game session in a queue (see CreateGameSessionQueue).
// When processing a placement request, Amazon GameLift searches for available
// resources on the queue's destinations, scanning each until it finds resources
// or the placement request times out.
//
// A game session placement request can also request player sessions. When a
// new game session is successfully created, Amazon GameLift creates a player
// session for each player included in the request.
//
// When placing a game session, by default Amazon GameLift tries each fleet
// in the order they are listed in the queue configuration. Ideally, a queue's
// destinations are listed in preference order.
//
// Alternatively, when requesting a game session with players, you can also
// provide latency data for each player in relevant regions. Latency data indicates
// the performance lag a player experiences when connected to a fleet in the
// region. Amazon GameLift uses latency data to reorder the list of destinations
// to place the game session in a region with minimal lag. If latency data is
// provided for multiple players, Amazon GameLift calculates each region's average
// lag for all players and reorders to get the best game play across all players.
//
// To place a new game session request, specify the following:
//
//    * The queue name and a set of game session properties and settings
//
//    * A unique ID (such as a UUID) for the placement. You use this ID to track
//    the status of the placement request
//
//    * (Optional) A set of player data and a unique player ID for each player
//    that you are joining to the new game session (player data is optional,
//    but if you include it, you must also provide a unique ID for each player)
//
//    * Latency data for all players (if you want to optimize game play for
//    the players)
//
// If successful, a new game session placement is created.
//
// To track the status of a placement request, call DescribeGameSessionPlacement
// and check the request's status. If the status is FULFILLED, a new game session
// has been created and a game session ARN and region are referenced. If the
// placement request times out, you can resubmit the request or retry it with
// a different queue.
//
//    * CreateGameSession
//
//    * DescribeGameSessions
//
//    * DescribeGameSessionDetails
//
//    * SearchGameSessions
//
//    * UpdateGameSession
//
//    * GetGameSessionLogUrl
//
//    * Game session placements StartGameSessionPlacement DescribeGameSessionPlacement
//    StopGameSessionPlacement
//
//    // Example sending a request using StartGameSessionPlacementRequest.
//    req := client.StartGameSessionPlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/StartGameSessionPlacement
func (c *Client) StartGameSessionPlacementRequest(input *StartGameSessionPlacementInput) StartGameSessionPlacementRequest {
	op := &aws.Operation{
		Name:       opStartGameSessionPlacement,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartGameSessionPlacementInput{}
	}

	req := c.newRequest(op, input, &StartGameSessionPlacementOutput{})
	return StartGameSessionPlacementRequest{Request: req, Input: input, Copy: c.StartGameSessionPlacementRequest}
}

// StartGameSessionPlacementRequest is the request type for the
// StartGameSessionPlacement API operation.
type StartGameSessionPlacementRequest struct {
	*aws.Request
	Input *StartGameSessionPlacementInput
	Copy  func(*StartGameSessionPlacementInput) StartGameSessionPlacementRequest
}

// Send marshals and sends the StartGameSessionPlacement API request.
func (r StartGameSessionPlacementRequest) Send(ctx context.Context) (*StartGameSessionPlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartGameSessionPlacementResponse{
		StartGameSessionPlacementOutput: r.Request.Data.(*StartGameSessionPlacementOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartGameSessionPlacementResponse is the response type for the
// StartGameSessionPlacement API operation.
type StartGameSessionPlacementResponse struct {
	*StartGameSessionPlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartGameSessionPlacement request.
func (r *StartGameSessionPlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}