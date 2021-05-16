package us.stephenmitchell.recipe_rest.model;

import javax.persistence.*;

@Entity
@Table(name = "step")
public class Step {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "ID")
    private Long id;

    @Column(name = "STEP_NUMBER")
    private Long step_number;

    @Column(name = "TEXT")
    private String text;

    @ManyToOne
    @JoinColumn(name = "RECIPE_ID")
    private Recipe recipe;

    public Step() { super(); }

    public Step(Long id, String text) {
        super();
        this.id = id;
        this.text = text;
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