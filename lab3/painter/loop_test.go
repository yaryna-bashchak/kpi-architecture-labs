package painter

import (
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
  "golang.org/x/exp/shiny/screen"
  "golang.org/x/image/draw"
  "image"
  "image/color"
  "testing"
  "time"
)

type MockReceiver struct {
  mock.Mock
}

func (r *MockReceiver) Update(t screen.Texture) {
  r.Called(t)
}

type MockScreen struct {
  mock.Mock
}

func (s *MockScreen) NewBuffer(size image.Point) (screen.Buffer, error) {
  return nil, nil
}

func (s *MockScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
  return nil, nil
}

func (s *MockScreen) NewTexture(size image.Point) (screen.Texture, error) {
  args := s.Called(size)
  return args.Get(0).(screen.Texture), args.Error(1)
}

type MockTexture struct {
  mock.Mock
}

func (t *MockTexture) Release() {
  t.Called()
}

func (t *MockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
  t.Called(dp, src, sr)
}

func (t *MockTexture) Bounds() image.Rectangle {
  args := t.Called()
  return args.Get(0).(image.Rectangle)
}

func (t *MockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
  t.Called(dr, src, op)
}

func (t *MockTexture) Size() image.Point {
  args := t.Called()
  return args.Get(0).(image.Point)
}

func TestLoop(t *testing.T) {
// Set up the MockReceiver, MockScreen, and MockTexture structs
screenMock := new(MockScreen)
textureMock := new(MockTexture)
receiverMock := new(MockReceiver)

// Create a texture and set up the expected behavior of the screen and receiver mocks
texture := image.Pt(400, 400)
screenMock.On("NewTexture", texture).Return(textureMock, nil)
receiverMock.On("Update", textureMock).Return()

// Create a painter loop and start it with the screenMock
loop := Loop{
  Receiver: receiverMock,
}
loop.Start(screenMock)

// Set up two mock operations and post them to the painter loop
op1 := new(MockOperation)
op2 := new(MockOperation)
textureMock.On("Bounds").Return(image.Rectangle{})
op1.On("Do", textureMock).Return(false)
op2.On("Do", textureMock).Return(true)
loop.Post(op1)
loop.Post(op2)

// Wait for a second and check that the queue is empty
time.Sleep(1 * time.Second)
assert.Equal(t, len(loop.MsgQueue.Queue), 0)

// Assert that the operations were called and the screen and receiver mocks were called with the expected arguments
op1.AssertCalled(t, "Do", textureMock)
op2.AssertCalled(t, "Do", textureMock)
receiverMock.AssertCalled(t, "Update", textureMock)
screenMock.AssertCalled(t, "NewTexture", image.Pt(400, 400))
}

type MockOperation struct {
  mock.Mock
}

func (o *MockOperation) Do(t screen.Texture) bool {
  args := o.Called(t)
  return args.Bool(0)
}
