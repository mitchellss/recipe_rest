package us.stephenmitchell.recipe_rest.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Table(name = "material")
public class Material {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "id")
    private Long id;

    @ManyToOne(cascade = CascadeType.DETACH)
    @JoinColumn(name = "recipe_id")
    private Recipe recipe;

    @Column(name = "material_number")
    private Long material_number;

    @Column(name = "measurement")
    private String measurement;

    @ManyToOne(cascade = CascadeType.DETACH)
    @JoinColumn(name = "ingredient")
    @Getter
    @Setter
    private Ingredient ingredient;

    @Column(name = "note")
    private String note;

    public Material() {
        super();
    }

    public Material(Recipe recipe, Long material_number, String measurement,
                    Ingredient ingredient, String note) {
        super();
        this.recipe = recipe;
        this.material_number = material_number;
        this.measurement = measurement;
        this.ingredient = ingredient;
        this.note = note;
    }

    public Long getId() {
        return id;
    }

    public Recipe getRecipe() {
        return recipe;
    }

    public void setRecipe(Recipe recipe) {
        this.recipe = recipe;
    }

    public Long getMaterial_number() {
        return material_number;
    }

    public void setMaterial_number(Long material_number) {
        this.material_number = material_number;
    }

    public String getMeasurement() {
        return measurement;
    }

    public void setMeasurement(String measurement) {
        this.measurement = measurement;
    }

    public String getNote() {
        return note;
    }

    public void setNote(String note) {
        this.note = note;
    }

    @Override
    public String toString() {
        return "Material{" +
                "id=" + id +
                ", recipe=" + recipe +
                ", material_number=" + material_number +
                ", measurement='" + measurement + '\'' +
                ", note='" + note + '\'' +
                '}';
    }
}
