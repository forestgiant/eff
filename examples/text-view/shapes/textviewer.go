package shapes

import (
	"log"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/tween"
	"github.com/forestgiant/eff/util"
)

const (
	text = `
In the near future, Detroit, Michigan is a dystopia and on the verge of total collapse due to financial ruin and a high crime rate. The mayor signs a deal with the mega-corporation Omni Consumer Products (OCP), giving it complete control of the underfunded Detroit Police Department. In exchange, OCP will be allowed to turn the run-down sections of Detroit into a high-end utopia called "Delta City", which will function as an independent city-state free of the United States.

OCP senior president Dick Jones proposes replacing the police with the ED-209 enforcement droid. At its first demonstration, however, ED-209 malfunctions and gruesomely kills a board member. Bob Morton, an ambitious employee, uses the opportunity to introduce his own experimental cyborg design, "RoboCop". To Jones's anger, the company chairman approves Morton's plan. Meanwhile, police officer Alex Murphy arrives at his new precinct following an OCP-directed transfer where he is introduced to his partner Anne Lewis. On their first patrol, they chase down a gang led by the ruthless criminal Clarence Boddicker, tailing them to an abandoned steel mill. When he and Lewis get separated, Murphy is caught and fatally injured by Boddicker's gang just before Boddicker himself executes the helpless cop. Morton selects Murphy for the RoboCop program and replaces most of his body with cybernetics, except for his brain and part of his digestive system.

RoboCop is given three primary directives: "Serve the public trust, Protect the innocent, and Uphold the law", as well as a classified fourth directive that Morton does not know of. He single-handedly and efficiently cleans Detroit of crime, earning Morton a promotion to vice president. Enraged, Jones hires Boddicker to murder Morton in his home. Meanwhile, Lewis realizes that RoboCop is really Murphy, and tells him his real name. RoboCop remembers past events from his life and returns to his former home, only to find that his wife and son have moved away. He connects to the police database, looks up Murphy's entry and discovers Boddicker's gang, who were responsible for his death.

RoboCop tracks down Boddicker to a cocaine factory and after a battle, threatens to kill him. Panicked, Boddicker admits his affiliation with Jones, verbally triggering RoboCop's law-abiding programming. RoboCop arrests Boddicker and turns him over to the police. He then confronts Jones and attempts to arrest him, but begins to shut down. Jones reveals that he planted the fourth directive, which prevents RoboCop from arresting any member of OCP's executive board. Jones explains his larger goal of taking over OCP, and confesses to Morton's murder before activating his personal ED-209 to destroy RoboCop. During the ensuing battle, Jones calls the police claiming that Robocop has malfunctioned and gone rogue. Robocop manages to escape ED-209, but is soon cornered by heavily armed police units and is nearly destroyed. Lewis helps RoboCop escape, and takes him to the abandoned steel mill. As RoboCop repairs himself, he and Lewis discuss his former life.

Under pressure by OCP and fearing their replacement by RoboCop, the police go on strike. Jones frees Boddicker and supplies his gang with anti-tank rifles and a tracking device to hunt down RoboCop. The gang converge on the steel mill, where RoboCop is able to kill most of them. Boddicker eventually subdues RoboCop, but dies after being stabbed in the throat before he can kill him.

RoboCop heads back to OCP headquarters, where Jones is presenting his improved ED-209 to the board. RoboCop plays a recording of Jones's confession, exposing his role in Morton's murder along with his sinister plans. Jones takes the chairman hostage, forcing a standoff as RoboCop explains that he cannot intervene due to the fourth directive. The chairman fires Jones from OCP, whereupon RoboCop shoots Jones, who falls through a window to his death far below. Grateful, the chairman says, "Nice shooting, son, what's your name?", to which RoboCop smiles and replies, "Murphy."
	`
)

// TextViewer Boilerplate drawable struct
type TextViewer struct {
	eff.Shape

	font       eff.Font
	textWidth  int
	textHeight int
	lines      []string
	tweener    tween.Tweener
}

// Init Boilerplate Init function, used to setup the drawable
func (t *TextViewer) Init(c eff.Canvas) {
	font, err := c.OpenFont("../assets/fonts/Jellee-Roman.ttf", 12)
	if err != nil {
		log.Fatal(err)
	}
	t.font = font
	minWidth := 5
	t.textWidth = minWidth
	t.lines, t.textHeight = util.GetMultilineText(t.font, text, t.textWidth, t.Graphics())
	t.tweener = tween.NewTweener(time.Second*5, func(progress float64) {
		maxWidth := c.Rect().W - minWidth
		t.textWidth = int(float64(maxWidth)*progress) + minWidth
		t.lines, t.textHeight = util.GetMultilineText(t.font, text, t.textWidth, t.Graphics())
	}, true, true, nil, nil)

	t.SetUpdateHandler(func() {
		t.tweener.Tween()

		t.Clear()
		linePoint := eff.Point{X: 0, Y: 0}
		for _, line := range t.lines {
			if linePoint.Y < c.Rect().H {
				t.DrawText(t.font, line, eff.Black(), linePoint)
			}

			linePoint.Y += t.textHeight
		}
	})
}
