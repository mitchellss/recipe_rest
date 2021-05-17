package us.stephenmitchell.recipe_rest.model;

import javax.persistence.*;

@Entity
@Table(name = "step")
public class Step {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "id")
    private Long id;

    @Column(name = "step_number")
    private Long step_number;

    @Column(name = "step_text")
    private String text;

    @ManyToOne(cascade = CascadeType.DETACH)
    @JoinColumn(name = "recipe_id")
    private Recipe recipe;

    public Step() { super(); }

    public Step(Long step_number, String text, Recipe recipe) {
        super();
        this.step_number = step_number;
        this.text = text;
        this.recipe = recipe;
    }

    public Long getId() {
        return id;
    }

    public Long getStep_number() {
        return step_number;
    }

    public void setStep_number(Long step_number) {
        this.step_number = step_number;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public Recipe getRecipe() {
        return recipe;
    }

    public void setRecipe(Recipe recipe) {
        this.recipe = recipe;
    }

    @Override
    public String toString() {
        return "StepModel{" +
                "id=" + id +
                ", step_number=" + step_number +
                ", recipe=" + recipe +
                '}';
    }
}
